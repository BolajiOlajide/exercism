//Solution goes in Sources
struct Year {
    let calendarYear: Int

    var isLeapYear: Bool {
        if calendarYear % 4 == 0 {
            if (calendarYear % 100 == 0) {
                if (calendarYear % 400 == 0) {
                    return true
                } else {
                    return false
                }
            } else {
                return true
            }
        } else {
            return false
        }
    }
}
