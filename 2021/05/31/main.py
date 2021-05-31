class Solution:
    def suggestedProducts(self, products: List[str], searchWord: str) -> List[List[str]]:
        products.sort()
        a = []
        def bs(q, d):
            l, h = 0, len(products)
            while l < h:
                m = l + (h - l) // 2
                p = products[m]
                if q <= p[:d]:
                    h = m - 1
                else:
                    l = m + 1
            return l
        for i in range(len(searchWord)):
            q = searchWord[:i+1]
            d = len(q)
            pos = bs(q, d)
            st = []
            while pos < len(products):
                p = products[pos]
                if q == p[:d]:
                    st.append(p)
                pos += 1
                if len(st) == 3:
                    break
            a.append(st)
        return a