package model



type Bubble struct {
	ID  int	`json:"id"`
	Title string `json:"title"`
	Status bool	`json:"status"`
}

func (b *Bubble)GetBubbles()(list []Bubble,err error){
		var todolist []Bubble
		err=DB.Debug().Find(&todolist).Error
		return todolist,err
}

func (b *Bubble)AddBubble()error{
	return DB.Debug().Create(b).Error
}

func(b *Bubble)UpdateBubble()error{
	return DB.Save(b).Error
}

func(b *Bubble)DeleteBubble(task string)error{
	err := DB.Where("id=?", task).First(b).Error
	if err!=nil{
		return err
	}
	return DB.Debug().Unscoped().Delete(b).Error
}