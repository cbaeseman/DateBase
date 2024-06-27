//
//  DateAppApp.swift
//  DateApp
//
//  Created by Chip on 6/26/24.
//

import SwiftUI

@main
struct DateAppApp: App {
    @State private var isActive: Bool = false
       
       var body: some Scene {
           WindowGroup {
               if isActive {
                   ContentView() // Replace with your main content view
               } else {
                   SplashScreen()
                       .onAppear {
                           DispatchQueue.main.asyncAfter(deadline: .now() + 2) {
                               withAnimation {
                                   self.isActive = true
                               }
                           }
                       }
               }
           }
       }
}
