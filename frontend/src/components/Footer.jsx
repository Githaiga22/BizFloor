const Footer = () => {
  return (
    <footer className="bg-gradient-to-b from-blue-900 to-gray-900">
      <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div className="col-span-1">
            <div className="flex items-center">
              <svg
                className="h-8 w-8 text-blue-700"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M13 10V3L4 14h7v7l9-11h-7z"
                />
              </svg>
              <span className="ml-2 text-xl font-bold text-white">BizFlow</span>
            </div>
            <p className="mt-4 text-gray-400">
              Empowering businesses with smart management solutions.
            </p>
          </div>
          <div className="col-span-1">
            <h3 className="text-white font-medium">Quick Links</h3>
            <ul className="mt-4 space-y-2">
              <li>
                <a href="/" className="text-gray-400 hover:text-blue-400">
                  Home
                </a>
              </li>
              <li>
                <a href="/about" className="text-gray-400 hover:text-blue-400">
                  About Us
                </a>
              </li>
              <li>
                <a href="/contact" className="text-gray-400 hover:text-blue-400">
                  Contact
                </a>
              </li>
            </ul>
          </div>
          <div className="col-span-1">
            <h3 className="text-white font-medium">Legal</h3>
            <ul className="mt-4 space-y-2">
              <li>
                <a href="/privacy" className="text-gray-400 hover:text-blue-400">
                  Privacy Policy
                </a>
              </li>
              <li>
                <a href="/terms" className="text-gray-400 hover:text-blue-400">
                  Terms of Service
                </a>
              </li>
            </ul>
          </div>
          <div className="col-span-1">
            <h3 className="text-white font-medium">Contact</h3>
            <ul className="mt-4 space-y-2">
              <li className="text-gray-400">support@bizflow.com</li>
              <li className="text-gray-400">+1 (555) 123-4567</li>
            </ul>
          </div>
        </div>
        <div className="mt-8 border-t border-gray-700 pt-8">
          <p className="text-center text-gray-400">
            2023 BizFlow. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
