package repository

import (
	"backend/model"
	"context"
)

func (t *RepositoryTest) TestRepository_ListRooms() {
	rooms, err := t.repository.ListRooms(context.Background(), model.RoomRequestParams{})
	t.Require().NoError(err)
	t.Len(rooms, 4)
}

func (t *RepositoryTest) TestRepository_ListRooms_occupied() {
	tr := true
	fl := false
	t.Run("occupied", func() {
		rooms, err := t.repository.ListRooms(context.Background(), model.RoomRequestParams{
			IsOccupied: &tr,
		})
		t.Require().NoError(err)
		t.Len(rooms, 2)
		t.Equal(room1.Number, rooms[0].Number)
		t.Equal(room2.Number, rooms[1].Number)
		for _, room := range rooms {
			t.NotEmpty(room)
		}
	})
	t.Run("unoccupied", func() {
		rooms, err := t.repository.ListRooms(context.Background(), model.RoomRequestParams{
			IsOccupied: &fl,
		})
		t.Require().NoError(err)
		t.Len(rooms, 2)
		t.Equal(room3, rooms[0])
		t.Equal(room4, rooms[1])
	})
}

func (t *RepositoryTest) TestRepository_ListRooms_Block() {
	rooms, err := t.repository.ListRooms(context.Background(), model.RoomRequestParams{
		Blocks: []string{"1"},
	})
	t.Require().NoError(err)
	t.Len(rooms, 1)
	t.Equal(room1.Number, rooms[0].Number)
}

func (t *RepositoryTest) TestRepository_GetRoom() {
	got, err := t.repository.GetRoom(context.Background(), room1.Number)
	t.Require().NoError(err)
	t.Equal(room1, got)
}
