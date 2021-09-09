package varible;

import "fmt";

type nameStr string;

func Name(name nameStr ) string {
    hallo := fmt.Sprintf("Hi , %v. Welcome!", name);
    return hallo;
};
