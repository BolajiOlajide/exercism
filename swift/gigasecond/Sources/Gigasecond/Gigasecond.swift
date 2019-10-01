//Solution goes in Sources
import Foundation

struct Gigasecond {
    let from: Date?

    init?(from stringDate: String) {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss"
        guard let from = dateFormatter.date(from: stringDate) else {
            return nil
        }
        self.from = dateFormatter.date(from: stringDate)
    }

    var description: String {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss"
        let newDate: Date = from!.addingTimeInterval(1000000000.0)
        return dateFormatter.string(from: newDate)
    }
}
