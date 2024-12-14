import React from 'react';
import Navbar from '../components/Navbar';
import Footer from '../components/Footer';

const Home = () => {
  const [showRegisterModal, setShowRegisterModal] = React.useState(false);

  return (
    <>
      <Navbar />
      <div className="min-h-screen flex flex-col">
        {/* Main Content */}
        <main className="flex-grow">
          {/* Hero Section */}
          <div className="bg-gradient-to-br from-blue-50 via-indigo-100 to-blue-200 py-20">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
              <div className="text-center">
                <h1 className="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl">
                  <span className="block">Streamline Your Business</span>
                  <span className="block text-blue-700">With BizFlow</span>
                </h1>
                <p className="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
                  The complete solution for online booking, payments, and record keeping. Manage your business efficiently with our comprehensive platform.
                </p>
                <div className="mt-5 max-w-md mx-auto sm:flex sm:justify-center md:mt-8">
                  <div className="rounded-md shadow">
                    <button
                      onClick={() => setShowRegisterModal(true)}
                      className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-700 hover:bg-blue-800 md:py-4 md:text-lg md:px-10"
                    >
                      Get Started
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Features Section */}
          <div className="py-12 bg-white">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
              <div className="grid grid-cols-1 gap-8 md:grid-cols-3">
                {/* Online Booking */}
                <div className="p-6 bg-white rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 hover:border-blue-200 border border-gray-100 hover:bg-blue-50">
                  <div className="text-blue-700 mb-4">
                    <svg className="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <h3 className="text-lg font-medium text-gray-900">Online Booking</h3>
                  <p className="mt-2 text-gray-500">
                    Effortlessly manage appointments and reservations with our intuitive booking system.
                  </p>
                </div>

                {/* Online Payment */}
                <div className="p-6 bg-white rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 hover:border-blue-200 border border-gray-100 hover:bg-blue-50">
                  <div className="text-blue-700 mb-4">
                    <svg className="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                    </svg>
                  </div>
                  <h3 className="text-lg font-medium text-gray-900">Secure Payments</h3>
                  <p className="mt-2 text-gray-500">
                    Process payments securely and efficiently with our integrated payment solutions.
                  </p>
                </div>

                {/* Record Keeping */}
                <div className="p-6 bg-white rounded-lg shadow-lg hover:shadow-xl transition-all duration-300 hover:border-blue-200 border border-gray-100 hover:bg-blue-50">
                  <div className="text-blue-700 mb-4">
                    <svg className="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <h3 className="text-lg font-medium text-gray-900">Record Keeping</h3>
                  <p className="mt-2 text-gray-500">
                    Keep track of all your business records in one centralized, secure location.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>

      {/* Modal for Registration */}
      {showRegisterModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
          <div className="bg-white p-6 rounded-lg shadow-lg w-96">
            <h2 className="text-2xl font-bold text-gray-900">Register</h2>
            <p className="mt-2 text-gray-600">Enter your details to register</p>
            {/* Add your registration form here */}
            <button
              onClick={() => setShowRegisterModal(false)}
              className="mt-4 px-4 py-2 bg-blue-700 text-white rounded"
            >
              Close
            </button>
          </div>
        </div>
      )}

      <Footer />
    </>
  );
};

export default Home;
