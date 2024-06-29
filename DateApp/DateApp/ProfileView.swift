//
//  ProfileView.swift
//  DateApp
//
//  Created by Clifford Baeseman on 6/26/24.
//

import SwiftUI

struct ProfileView: View {
    var body: some View {
        VStack(alignment: .center) {
            // Profile Picture
            Image("profile_picture") // Replace with the actual image name or URL
                .resizable()
                .aspectRatio(contentMode: .fill)
                .frame(width: 150, height: 150)
                .clipShape(Circle())
                .overlay(Circle().stroke(Color.white, lineWidth: 4))
                .shadow(radius: 10)
            
            // Name and Age
            Text("Cliff Baeseman, 38")
                .font(.largeTitle)
                .fontWeight(.bold)
                .padding(.top, 10)
            
            // Location
            Text("Vero Beach, FL")
                .font(.subheadline)
                .foregroundColor(.gray)
                .padding(.bottom, 10)
            
            // Tagline
            Text("\"Adventurous and fun-loving!\"")
                .font(.headline)
                .italic()
                .padding(.bottom, 20)
            
            // Profile Description
            Text("Hi! I'm Cliff, a software developer who loves exploring new places and trying out new cuisines. In my free time, I enjoy hiking, reading, and spending time with my friends and family. Looking forward to meeting someone who shares similar interests and is up for an adventure!")
                .font(.body)
                .multilineTextAlignment(.center)
                .padding(.horizontal, 40)
            
            Spacer()
        }
        .padding()
        .background(Color(UIColor.systemBackground))
        .edgesIgnoringSafeArea(.all)
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}

#Preview {
    ProfileView()
}
