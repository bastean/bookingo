package valueobj

import (
	"regexp"
	"strings"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

const RoomMinCharactersLength = "1"
const RoomMaxCharactersLength = "64"

type Room struct {
	Room string
}

func (room *Room) Value() string {
	return room.Room
}

func (room *Room) IsValid() error {
	validate := regexp.MustCompile(`^[\w\s-]{1,64}$`)

	if !validate.MatchString(room.Room) {
		return errors.Default()
	}

	return nil
}

func NewRoom(room string) (models.ValueObject[string], error) {
	room = strings.TrimSpace(room)

	roomVO := &Room{
		Room: room,
	}

	if roomVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewRoom",
			What:  "room must be between " + RoomMinCharactersLength + " to " + RoomMaxCharactersLength + " characters",
			Why: errors.Meta{
				"Room": room,
			},
		})
	}

	return roomVO, nil
}
