# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: adjoly <adjoly@student.42angouleme.fr>     +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2024/12/01 20:34:14 by adjoly            #+#    #+#              #
#    Updated: 2024/12/01 20:34:38 by adjoly           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

MODNAME = github.com/keyzox71/timmy

all: buildcli

tidy:
	@go mod tidy


clean:
	@rm -f $(CLI_NAME)
	@printf "$(RED)「🗑️」($(CLI_NAME)) Program deleted\n"

fclean: clean

re: fclean all

.PHONY: re fclean clean tidy buildcli
