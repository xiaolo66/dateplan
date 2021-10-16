package servers

import (
	"dateplan/dto"
	"dateplan/model"
)

type Bubble struct {

}

func NewBubble()*Bubble{
	return &Bubble{}
}

func(b *Bubble)GetAllBubble()(list []model.Bubble,err error){
	var s model.Bubble
	return s.GetBubbles()
}

func(b *Bubble)AddBubble(info dto.Bubble)error{
	var s model.Bubble
	s.Title=info.Title
	s.Title=info.Title
	return s.AddBubble()
}

func(b *Bubble)UpdateBubble(id string)error{
	var s model.Bubble
	if err:=model.DB.Debug().Where("id=?",id).First(&s).Error;err!=nil{
		return err
	}
	if s.Status==true{
		s.Status=false
	}else {
		s.Status=true
	}
	return s.UpdateBubble()
}

func(b *Bubble)DeleteBubble(id string)error{
	var s model.Bubble
	return s.DeleteBubble(id)
}


