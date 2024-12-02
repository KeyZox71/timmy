# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2024/12/01 20:34:14 by adjoly            #+#    #+#              #
#    Updated: 2024/12/02 15:57:16 by adjoly           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

MOD_NAME = github.com/keyzox71/timmy

EXEC_NAME = timmy

RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
PURPLE = \e[0;35m
NC = \033[0m
DELETE = \x1B[2K\r

all: build-timmy

tidy:
	@go mod tidy

build-timmy:
	@go build -o $(EXEC_NAME) $(MOD_NAME)/cmd/timmy
	@printf "$(YELLOW)「✨」($(EXEC_NAME)) Program compiled\n"

clean:
	@rm -f $(EXEC_NAME)
	@printf "$(RED)「🗑️」($(EXEC_NAME)) Program deleted\n"

fclean: clean

re: fclean all

.PHONY: re fclean clean tidy buildcli
