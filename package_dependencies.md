# Dependency Table For: github.com/flowdev/example-mono/

| | b a c k e n d / c a t a l o g u e - D | b a c k e n d / r e c o m m e n d a t i o n s / b u y h i s t - S | b a c k e n d / r e c o m m e n d a t i o n s / e x t e r n a l a c t s - S | b a c k e n d / r e c o m m e n d a t i o n s / i n t e r n a l a c t s - S | b a c k e n d / u s e r - S | e m a i l - S | e n t i t i e s - T | f r o n t e n d / c u s t o m e r - D | f r o n t e n d / p a y m e n t - S | f r o n t e n d / p a y m e n t / b a n k t r a n s f e r - S | f r o n t e n d / p a y m e n t / c r e d i t c a r d - S | f r o n t e n d / p a y m e n t / p a y p a l - S | f r o n t e n d / p a y m e n t / v o u c h e r s - S | f r o n t e n d / s h o p p i n g c a r t - S | x / l o g g i n g - T |
| :- | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| `backend/catalogue` | | | | | | | `T` | | | | | | | | `T` |
| **backend/cmd/back4front** | **D** | **S** | **S** | **S** | | | **T** | | | | | | | | **T** |
| **backend/cmd/shopback** | **D** | **S** | **S** | **S** | **S** | | **T** | | | | | | | | **T** |
| backend/recommendations/buyhist | D | | | | | | T | | | | | | | | T |
| backend/recommendations/externalacts | | | | | | | | D | | | | | | | T |
| backend/recommendations/internalacts | D | | | | | | T | D | | | | | | | T |
| backend/stats/custstat | | | | | | | | D | | | | | | | T |
| backend/stats/prodstat | D | | | | | | T | | | | | | | | T |
| backend/user | | | | | | S | | | | | | | | | T |
| email | | | | | | | | | | | | | | | T |
| **frontend/cmd/shopfront** | | | | | | | | **D** | **S** | | | | | **S** | **T** |
| `frontend/customer` | | | | | | `S` | | | | | | | | | `T` |
| frontend/payment | D | | | | | S | T | D | | S | S | S | S | | T |
| frontend/payment/banktransfer | | | | | | | T | | | | | | | | T |
| frontend/payment/creditcard | | | | | | | T | | | | | | | | T |
| frontend/payment/paypal | | | | | | | T | | | | | | | | T |
| frontend/payment/vouchers | | | | | | | T | | | | | | | | T |
| frontend/shoppingcart | D | | | | | | T | | | | | | | | T |

### Legend

* Rows - Importing packages
* columns - Imported packages


#### Meaning Of Row And Row Header Formating

* **Bold** - God package
* `Code` - Database package
* _Italic_ - Tool package


#### Meaning Of Letters In Table Columns

* G - God package
* D - Database package
* T - Tool package
* S - Standard package
