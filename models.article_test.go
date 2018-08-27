package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	aList := getAllArticles()

	if len(aList) != len(articleList) {
		t.Fail()
	}

	for i,v := range aList {
		if v.Content != articleList[i].Content ||
			v.ID    !=  articleList[i].ID ||
			v.Title !=  articleList[i].Title {
				t.Fail()
				break
		}
	}
}


func TestGetArticleByID(t *testing.T) {
	a, err := getArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}