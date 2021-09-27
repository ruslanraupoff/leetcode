class Solution:
    def numUniqueEmails(self, emails: List[str]) -> int:
        h = {}
    
        def clean(email):
            i = email.index('@')
            user, atDomain = email[:i], email[i:]
            return deleteDot(trim(user)) + atDomain
        
        def deleteDot(username):
            return username[::-1].replace(".", "")

        def trim(username):
            i = username.find('+')
            if i == -1:
                return username
            return username[:i]
            
        for email in emails:
            h[clean(email)] = True
            
        return len(h)

