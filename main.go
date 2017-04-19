package main

import (
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

var stretches = []struct {
	Title       string
	Description string
}{
	// sourced from http://www.besthealthmag.ca/best-you/stretching/8-stretches-you-can-do-at-your-desk
	{"Shrug your shoulders", "Inhale deeply and shrug your shoulders, lifting them high up to your ears. Hold. Release and drop."},
	{"Seated Glute Stretch", "Sit on a chair; bring one ankle up onto the knee of your other leg. Lean forward, keeping your back straight, and you should immediately feel the stretch in the gluteus maximus. Hold for 16 seconds. Continue forward as far as you can comfortably go, pushing gently on the top knee with your hand. Hold for 30 seconds, then switch sides."},
	{"Upper Trapezius Stretch", "Sitting up straight, with your feet flat and shoulders back, grab onto the bottom of your chair with your right hand. Then slowly tilt your head sideways, bringing your left ear toward your left shoulder until you feel a gentle stretch along the right side of your neck and shoulder. Hold for 10 to 15 seconds; repeat on the opposite side. Do this stretch a few times a day to help prevent a strain from starting in the first place.."},
	{"Wrist Flexor Stretch", "While seated, extend your right arm forward at shoulder height. Keeping your elbow straight, grasp your right hand with your left, and slowly bend the wrist backward until you feel a stretch along the bottom of your forearm. Hold for 15 seconds. Then bend wrist downward until a stretch is felt on the top of the arm, and hold for 15 seconds. Switch arms and repeat. Repeat four times for each hand."},
	{"Hamstring stretch with chair", "Position yourself about a foot and a half (half a metre) behind a chair. Stand up straight, with your shoulders back and feet pointing forward. Keep your knees straight, but not locked. Keeping your neck aligned with your back throughout this move, bend at the hips to a 90-degree angle. (You may need to move the chair forward so you are putting the weight of your body on the chair and not holding yourself up-you’ll feel the stretch more this way.) Hold for 30 seconds."},
	{"Prayer Stretch", "Place palms together, fingers straight, in a prayer position, elbows pointed out to your sides. Keeping the heels of the hands together, slowly lower hands and raise elbows so the angle at the wrist decreases. Press palms and fingers together firmly for five seconds, then release. Repeat five times."},
	{"Hip Stretch", "Sitting at your desk, cross your right foot over your left knee, extending the left leg out a little if necessary. Keeping your right foot flexed to protect the ankle, lift the upper body tall, then hinge forward at the hips with a flat back. Hold the stretch for several deep breaths, lengthening the outer hip muscles. Repeat on the other side.."},
}

func main() {

	for {
		for time.Now().Minute() == 15 || time.Now().Minute() == 45 {
			// Remind to get out of the chair every time
			Stretch("1/2. Stand up and sit down", "Get some water, use the restroom, or go for a walk")
			// shift
			s, stretches := stretches[0], stretches[1:]
			Stretch("2/2. "+s.Title, s.Description)
			// push
			stretches = append(stretches, s)
			time.Sleep(15 * time.Minute)
		}
	}
}

func Stretch(title, message string) {
	notify = notificator.New(notificator.Options{
		AppName: "Stretch Reminder",
	})

	//TODO: add icons showing the movements
	notify.Push(title, message, "", notificator.UR_CRITICAL)
}
