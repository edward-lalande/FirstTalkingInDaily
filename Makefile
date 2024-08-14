
NAME	=	FirstTalkingInDaily
GO	=	go

all:
	$(GO) build -o $(NAME) ./src/*.go

clean:
	$(RM) $(NAME)

re:	clean all

.PHONY	=	all clean all
