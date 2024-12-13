import React from 'react';

const Navbar = ({ isOpen, setIsOpen, setShowRegisterModal }) => {
  return (
    <nav className="bg-gradient-to-r from-blue-50 to-indigo-100 shadow-lg">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <div className="flex-shrink-0 flex items-center">
              <svg className="h-8 w-8 text-blue-700" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              <span className="ml-2 text-xl font-bold text-gray-800">BizFlow</span>
            </div>
          </div>
          
          {/* Desktop Menu */}
          <div className="hidden md:flex md:items-center md:space-x-6">
            <a href="/" className="text-gray-600 hover:text-blue-400">Home</a>
            <a href="/about" className="text-gray-600 hover:text-blue-400">About Us</a>
            <a href="/contact" className="text-gray-600 hover:text-blue-400">Contact</a>
            <a href="/login" className="text-gray-600 hover:text-blue-400">Login</a>
            <button 
              onClick={() => setShowRegisterModal(true)}
              className="bg-blue-700 text-white px-4 py-2 rounded-md hover:bg-blue-800 transition-colors duration-300"
            >
              Register
            </button>
          </div>

          {/* Mobile menu button */}
          <div className="md:hidden flex items-center">
            <button
              onClick={() => setIsOpen(!isOpen)}
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
            >
              <svg
                className={`${isOpen ? 'hidden' : 'block'} h-6 w-6`}
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
              <svg
                className={`${isOpen ? 'block' : 'hidden'} h-6 w-6`}
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      {/* Mobile Menu */}
      <div className={`${isOpen ? 'block' : 'hidden'} md:hidden`}>
        <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
          <a href="/" className="block px-3 py-2 text-gray-600 hover:text-blue-400">Home</a>
          <a href="/about" className="block px-3 py-2 text-gray-600 hover:text-blue-400">About Us</a>
          <a href="/contact" className="block px-3 py-2 text-gray-600 hover:text-blue-400">Contact</a>
          <a href="/login" className="block px-3 py-2 text-gray-600 hover:text-blue-400">Login</a>
          <button 
            onClick={() => setShowRegisterModal(true)}
            className="block w-full text-left px-3 py-2 text-gray-600 hover:text-blue-400"
          >
            Register
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;