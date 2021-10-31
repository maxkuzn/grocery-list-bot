package answer

import (
	"fmt"

	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

const (
	CannotRespond = "Я не могу ответить на предыдущее сообщение :("
	internalError = "Произошла внутренняя ошибка. Приношу свои извинения!\n%v"

	HelpSuggestion  = "Посмотреть список команд /help."
	OnlyCommands    = "Пока что я умею отвечать только на команды.\n" + HelpSuggestion
	UnknownCommands = "Я не знаю такой команды.\n" + HelpSuggestion
	NotImplemented  = "Пока что это я не умею, но обязательно научусь в будущем!"

	HelpHelp    = "/help - Вывести это сообщение"
	CreateHelp  = "/create <имя списка> - создать список с именем <имя списка>"
	SwitchHelp  = "/switch <имя списка> - переключиться на список с именем <имя списка>"
	DeleteHelp  = "/delete <имя списка> - удалить список с именем <имя списка>"
	AddHelp     = "/add <запись> - добавить запись в список"
	RemoveHelp  = "/remove <N> - удалить из списка запись под номером <N>"
	CheckHelp   = "/check <N> - отметить запись под номером <N>"
	UncheckHelp = "/uncheck <N> - снять отметку c запись под номером <N>"
	ShowHelp    = "/show - показать список"
	Help        = "Список команд:\n" +
		HelpHelp + "\n" +
		"\n" +
		CreateHelp + "\n" +
		SwitchHelp + "\n" +
		DeleteHelp + "\n" +
		"\n" +
		AddHelp + "\n" +
		RemoveHelp + "\n" +
		CheckHelp + "\n" +
		UncheckHelp + "\n" +
		ShowHelp + "\n"

	listCreated  = "Cписок %q [id=%d] успешно создан"
	listDeleted  = "Cписок %q [id=%d] успешно удален"
	listSelected = "Выбран список %q [id=%d]"

	listNameExists  = "Cписок c именем %q уже существует"
	listDoesntExist = "Cписока c именем %q не существует"
)

func InternalError(err error) string {
	return fmt.Sprintf(internalError, err)
}

func ListCreated(listName string, listID model.ListID) string {
	return fmt.Sprintf(listCreated, listName, listID)
}

func ListDeleted(listName string, listID model.ListID) string {
	return fmt.Sprintf(listDeleted, listName, listID)
}

func ListSelected(listName string, listID model.ListID) string {
	return fmt.Sprintf(listSelected, listName, listID)
}

func ListNameExists(listName string) string {
	return fmt.Sprintf(listNameExists, listName)
}

func ListDoesntExist(listName string) string {
	return fmt.Sprintf(listDoesntExist, listName)
}
