package gateway

import (
	"errors"
	"github.com/Goscord/goscord/discord"
	"sync"
)

type State struct {
	session  *Session
	mut      *sync.RWMutex
	Guilds   map[string]*discord.Guild
	Channels map[string]*discord.Channel
	Members  map[string]map[string]*discord.Member
}

func NewState(session *Session) *State {
	return &State{
		session:  session,
		mut:      &sync.RWMutex{},
		Guilds:   map[string]*discord.Guild{},
		Channels: map[string]*discord.Channel{},
		Members:  map[string]map[string]*discord.Member{},
	}
}

// GUILDS

func (s *State) AddGuild(guild *discord.Guild) {
	// TODO : Members

	if guild.Channels != nil {
		for _, channel := range guild.Channels {
			s.AddChannel(channel)
		}
	}

	s.mut.Lock()
	s.Guilds[guild.Id] = guild
	s.mut.Unlock()
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	if g, err := s.Guild(guild.Id); err == nil {
		if guild.Channels != nil {
			for _, channel := range guild.Channels {
				s.RemoveChannel(channel)
			}
		}

		s.mut.Lock()
		delete(s.Guilds, g.Id)
		s.mut.Unlock()
	}
}

func (s *State) Guild(id string) (*discord.Guild, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	if guild, ok := s.Guilds[id]; ok {
		return guild, nil
	}

	return nil, errors.New("Guild not found")
}

// CHANNELS

func (s *State) AddChannel(channel *discord.Channel) {
	s.mut.Lock()
	s.Channels[channel.Id] = channel
	s.mut.Unlock()
}

func (s *State) RemoveChannel(channel *discord.Channel) {
	if c, err := s.Channel(channel.Id); err == nil {
		s.mut.Lock()
		delete(s.Channels, c.Id)
		s.mut.Unlock()
	}
}

func (s *State) Channel(id string) (*discord.Channel, error) {
	s.mut.RLock()

	if channel, ok := s.Channels[id]; ok {
		s.mut.RUnlock()
		return channel, nil
	}

	s.mut.RUnlock()

	channel, _ := s.session.Channel.GetChannel(id)

	if channel != nil {
		s.AddChannel(channel)

		return channel, nil
	}

	return nil, errors.New("Channel not found")
}

// MEMBERS

func (s *State) AddMember(guildID string, member *discord.Member) {
	if _, err := s.Guild(guildID); err != nil {
		return
	}

	s.mut.Lock()
	if _, ok := s.Members[guildID]; !ok {
		s.Members[guildID] = map[string]*discord.Member{}
	}
	s.Members[guildID][member.User.Id] = member
	s.mut.Unlock()
}

func (s *State) RemoveMember(guildID string, member *discord.Member) {
	if _, err := s.Guild(guildID); err != nil {
		return
	}

	s.mut.Lock()
	if _, ok := s.Members[guildID]; ok {
		delete(s.Members[guildID], member.User.Id)
	}
	s.mut.Unlock()
}

func (s *State) Member(guildID string, userID string) (*discord.Member, error) {
	if _, err := s.Guild(guildID); err != nil {
		return nil, err
	}

	s.mut.RLock()
	if _, ok := s.Members[guildID]; ok {
		if member, ok := s.Members[guildID][userID]; ok {
			s.mut.RUnlock()
			return member, nil
		}
	}
	s.mut.RUnlock()

	member, _ := s.session.Guild.GetMember(guildID, userID)

	if member != nil {
		s.AddMember(guildID, member)

		return member, nil
	}

	return nil, errors.New("Member not found")
}
