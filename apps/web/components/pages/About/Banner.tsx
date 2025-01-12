import type { FC } from 'react'
import Link from 'next/link'

export const Banner: FC = () => {
  return (
    <section className="py-12 md:py-16 border-b border-zinc-300 border-solid dark:bg-zinc-900 dark:border-zinc-600">
      <div className="md:max-w-4xl lg:max-w-7xl text-center ml-auto mr-auto w-11/12">
        <h1 className="font-bold text-4xl md:text-6xl mb-6 text-black dark:text-white">
          Universal Micro Services
        </h1>
        <h2 className="text-md md:text-lg max-w-xl mx-auto text-zinc-700 dark:text-zinc-400 font-medium">
          Explore, discover and consume public APIs as simpler building blocks
        </h2>
        <div className="mt-6">
          <Link href="/pricing">
            <a className="inline-flex items-center btn">
              Get Started
            </a>
          </Link>
        </div>
        <p className="mt-6 text-sm">
          Powered by{' '}
          <a href="https://micro.dev" target="_blank" rel="noreferrer">
            Micro
          </a>
        </p>
      </div>
    </section>
  )
}
