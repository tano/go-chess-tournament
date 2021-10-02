package core

var participants []string

func GetParticipants() []string {
	return participants
}

func AddParticipant(participant string) {
	participants = append(participants, participant)
}

func RemoveParticipant(participant string) {
	targetIndex := -1
	for i := 0; i < len(participants); i++ {
		if participants[i] == participant {
			targetIndex = i
		}
	}

	participants = append(participants[:targetIndex], participants[targetIndex:]...)
}
