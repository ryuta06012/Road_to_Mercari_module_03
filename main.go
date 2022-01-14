/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/11 07:11:56 by hryuuta           #+#    #+#             */
/*   Updated: 2022/01/14 13:28:59 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments.")
	}
	http.HandleFunc("/ping", fortune.pingHandler)
	if err := http.ListenAndServe(":"+os.Args[1], nil); err != nil {
		log.Fatal(err)
	}
}
