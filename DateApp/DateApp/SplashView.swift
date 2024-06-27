//
//  SplashView.swift
//  DateApp
//
//  Created by Chip on 6/26/24.
//

import SwiftUI

struct SplashScreen: View {
    @State private var opacity: Double = 0.0
    
    var body: some View {
        ZStack {
            Color.white
                .edgesIgnoringSafeArea(.all)
            Image("datebase.png")
                .resizable()
                .scaledToFit()
                .frame(width: 200, height: 200)
                .opacity(opacity)
                .onAppear {
                    withAnimation(.easeIn(duration: 1.0)) {
                        opacity = 1.0
                    }
                }
        }
    }
}

#Preview {
    SplashScreen()
}
