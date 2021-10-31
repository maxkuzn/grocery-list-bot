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
	itemAdded    = "Запись добавлена в список"
	itemDeleted  = "Запись удалена из списока"

	showEmptyList  = "Список %q пустой"
	showListHeader = "Список %q:"
	showListItem   = "%d. %s"

	listNameExists  = "Cписок c именем %q уже существует"
	listDoesntExist = "Cписока c именем %q не существует"
	ListNotSelected = "Ни один список не выбран\nВыберете список командой /switch или создайте командой /create\nПодробнее: /help"
	InvalidIndex    = "Записи с таким номером не существует\nПосмотреть актуальный список /show"
)

// Success
func ListCreated(listName string, listID model.ListID) string {
	return fmt.Sprintf(listCreated, listName, listID)
}

func ListDeleted(listName string, listID model.ListID) string {
	return fmt.Sprintf(listDeleted, listName, listID)
}

func ListSelected(listName string, listID model.ListID) string {
	return fmt.Sprintf(listSelected, listName, listID)
}

func ItemAdded() string {
	return fmt.Sprintf(itemAdded)
}

func ItemDeleted() string {
	return fmt.Sprintf(itemDeleted)
}

func ShowList(list model.List) string {
	if len(list.Items) == 0 {
		return fmt.Sprintf(showEmptyList, list.Name)
	}
	text := fmt.Sprintf(showListHeader, list.Name) + "\n"
	items := list.GetOrderedItems()
	for i, item := range items {
		text += fmt.Sprintf(showListItem, i+1, item.Description) + "\n"
	}
	return text
}

// Errors
func InternalError(err error) string {
	return fmt.Sprintf(internalError, err)
}

func ListNameExists(listName string) string {
	return fmt.Sprintf(listNameExists, listName)
}

func ListDoesntExist(listName string) string {
	return fmt.Sprintf(listDoesntExist, listName)
}
