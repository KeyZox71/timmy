# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2024/12/01 20:34:14 by adjoly            #+#    #+#              #
#    Updated: 2024/12/02 14:46:34 by adjoly           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

MODNAME = github.com/keyzox71/timmy

all: build-timmy

tidy:
	@go mod tidy

build-timmy:
	@go build $(MODNAME)/cmd/timmy

build-tm:
	@go build $(MODNAME)/cmd/tm

build-ts:
	@go build $(MODNAME)/cmd/ts

clean:
	@rm -f $(CLI_NAME)
	@printf "$(RED)「🗑️」($(CLI_NAME)) Program deleted\n"

fclean: clean

re: fclean all

.PHONY: re fclean clean tidy buildcli
