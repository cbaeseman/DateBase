//
//  MessagesView.swift
//  DateApp
//
//  Created by Clifford Baeseman on 6/26/24.
//

import SwiftUI

struct Message: Identifiable {
    let id = UUID()
    let thumbnail: String
    let name: String
    let message: String
}

struct MessageRow: View {
    var message: Message

    var body: some View {
        HStack {
            Image(message.thumbnail)
                .resizable()
                .frame(width: 50, height: 50)
                .clipShape(Circle())
            VStack(alignment: .leading) {
                Text(message.name)
                    .font(.headline)
                Text(message.message)
                    .font(.subheadline)
                    .lineLimit(1)
            }
        }
        .padding(.vertical, 5)
    }
}


struct MessagesView: View {
    let messages = [
        Message(thumbnail: "user1", name: "Alice", message: "Hey, how are you doing today?"),
        Message(thumbnail: "user2", name: "Bob", message: "Are you free this weekend?"),
        Message(thumbnail: "user3", name: "Catherine", message: "Let's catch up soon! It's been a while."),
    ]

    var body: some View {
        NavigationView {
            List(messages) { message in
                MessageRow(message: message)
            }
            .navigationBarTitle("Messages")
        }
    }
}

#Preview {
    MessagesView()
}
