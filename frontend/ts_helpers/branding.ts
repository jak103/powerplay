declare const __brand: unique symbol
type Brand<B> = { [__brand]: B }
/**
 * A branded type is the same as an unbranded type at runtime.
 * 
 * At compile time, functions that accept a branded type as a parameter 
 * will not accept an unbranded type. This improves type safety.
 * 
 * @template T The type to be used at runtime. `Branded<T,_>` will act exactly like `T`
 * @template B A *unique* string that identifies the particular brand of this type
 * 
 * @example
 * ```ts
 * type EmailAddress = Branded<string, 'email_address'>
 * 
 * const createEmailAddress = (s: string) => s as EmailAddress
 * const sendEmail = (to: EmailAddress) => {}
 * const logString = (s: string) => console.log(s)
 * 
 * sendEmail('xyz') // compile error
 * sendEmail('john@example.com') // compile error
 * sendEmail(createEmailAddress('alice@example.com')) // okay!
 * logString(createEmailAddress('bob@example.com')) // prints 'bob@example.com'
 * ```
 * 
 * @see https://egghead.io/blog/using-branded-types-in-typescript
 */
export type Branded<T, B> = T & Brand<B>
