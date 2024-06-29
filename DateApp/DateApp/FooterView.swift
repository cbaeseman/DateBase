//
//  FooterView.swift
//  DateApp
//
//  Created by Clifford Baeseman on 6/28/24.
//

import SwiftUI

struct FooterView: View {
    var body: some View {
        HStack(spacing:0) {
            Button(action: {}) {
                Image("refresh")
            }
            Button(action: {}) {
                Image("dismiss")
            }
            Button(action: {}) {
                Image("super")
            }
            Button(action: {}) {
                Image("like")
            }
            Button(action: {}) {
                Image("boost")
            }
        }
    }
}

struct FooterSection_Previews: PreviewProvider {
    static var previews: some View {
        FooterView()
    }
}
#Preview {
    FooterView()
}
