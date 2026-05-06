import React, { useState } from 'react';
import { 
  CloudArrowUpIcon, 
  DocumentIcon, 
  DocumentTextIcon,
  CheckCircleIcon,
  ArrowPathRoundedSquareIcon,
  ArrowDownTrayIcon,
  MoonIcon,
  GlobeAltIcon,
  ClockIcon,
  DocumentChartBarIcon
} from '@heroicons/react/24/outline';

export default function DocuCompareDashboard() {
  return (
    <div className="min-h-screen bg-[#F8FAFC] text-slate-800 font-sans">
      {/* Top Navbar */}
      <header className="bg-[#1E1B4B] text-white flex justify-between items-center px-6 py-4 shadow-md">
        <div className="flex items-center space-x-3">
          <div className="bg-indigo-500 p-2 rounded-lg">
            <DocumentChartBarIcon className="h-6 w-6 text-white" />
          </div>
          <div>
            <h1 className="text-xl font-bold tracking-wide">DocuCompare</h1>
            <p className="text-xs text-indigo-200">Image/PDF vs DOCX Checker</p>
          </div>
        </div>
        
        <div className="flex items-center space-x-4">
          <button className="flex items-center space-x-2 bg-indigo-900/50 hover:bg-indigo-800 border border-indigo-700 rounded-lg px-4 py-2 text-sm transition">
            <GlobeAltIcon className="h-5 w-5 text-indigo-300" />
            <span>বাংলা (Bangla)</span>
          </button>
          <button className="p-2 rounded-full hover:bg-indigo-800 transition">
            <MoonIcon className="h-5 w-5 text-indigo-200" />
          </button>
          <button className="flex items-center space-x-2 bg-indigo-500 hover:bg-indigo-600 rounded-lg px-4 py-2 text-sm font-medium transition shadow-sm">
            <ClockIcon className="h-5 w-5" />
            <span>Compare History</span>
          </button>
        </div>
      </header>

      <main className="max-w-7xl mx-auto p-6 space-y-6">
        {/* Upload Section */}
        <section className="bg-white rounded-xl shadow-sm border border-slate-200 p-6 grid grid-cols-1 md:grid-cols-3 gap-6">
          
          {/* Step 1: First Document */}
          <div className="space-y-4">
            <div className="flex items-center space-x-3">
              <span className="flex items-center justify-center w-8 h-8 rounded-full bg-indigo-100 text-indigo-600 font-bold text-sm">1</span>
              <div>
                <h3 className="font-semibold text-slate-800">Upload First Document</h3>
                <p className="text-xs text-slate-500">Image or PDF file</p>
              </div>
            </div>
            <div className="border-2 border-dashed border-indigo-200 rounded-xl bg-indigo-50/50 flex flex-col items-center justify-center py-10 px-4 text-center hover:bg-indigo-50 transition cursor-pointer">
              <DocumentIcon className="h-10 w-10 text-indigo-500 mb-3" />
              <p className="text-sm text-slate-600 mb-1">Drag & drop image or PDF here</p>
              <p className="text-xs text-slate-400 mb-4">or</p>
              <button className="bg-indigo-500 hover:bg-indigo-600 text-white px-5 py-2 rounded-lg text-sm font-medium transition">
                Choose File
              </button>
            </div>
            {/* File uploaded state */}
            <div className="flex items-center justify-between p-3 bg-slate-50 border border-slate-200 rounded-lg">
              <div className="flex items-center space-x-3">
                <div className="bg-red-100 p-1.5 rounded text-red-600 text-xs font-bold">PDF</div>
                <div>
                  <p className="text-sm font-medium text-slate-700">document-image.pdf</p>
                  <p className="text-xs text-slate-500">2.45 MB</p>
                </div>
              </div>
              <CheckCircleIcon className="h-6 w-6 text-emerald-500" />
            </div>
          </div>

          {/* Step 2: Second Document */}
          <div className="space-y-4">
            <div className="flex items-center space-x-3">
              <span className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-100 text-blue-600 font-bold text-sm">2</span>
              <div>
                <h3 className="font-semibold text-slate-800">Upload Second Document</h3>
                <p className="text-xs text-slate-500">DOCX or PDF file</p>
              </div>
            </div>
            <div className="border-2 border-dashed border-blue-200 rounded-xl bg-blue-50/50 flex flex-col items-center justify-center py-10 px-4 text-center hover:bg-blue-50 transition cursor-pointer">
              <DocumentTextIcon className="h-10 w-10 text-blue-500 mb-3" />
              <p className="text-sm text-slate-600 mb-1">Drag & drop DOCX or PDF here</p>
              <p className="text-xs text-slate-400 mb-4">or</p>
              <button className="bg-blue-500 hover:bg-blue-600 text-white px-5 py-2 rounded-lg text-sm font-medium transition">
                Choose File
              </button>
            </div>
            {/* File uploaded state */}
            <div className="flex items-center justify-between p-3 bg-slate-50 border border-slate-200 rounded-lg">
              <div className="flex items-center space-x-3">
                <div className="bg-blue-100 p-1.5 rounded text-blue-600 text-xs font-bold">DOCX</div>
                <div>
                  <p className="text-sm font-medium text-slate-700">document-file.docx</p>
                  <p className="text-xs text-slate-500">1.32 MB</p>
                </div>
              </div>
              <CheckCircleIcon className="h-6 w-6 text-emerald-500" />
            </div>
          </div>

          {/* Step 3: Compare Settings */}
          <div className="space-y-4">
            <div className="flex items-center space-x-3">
              <span className="flex items-center justify-center w-8 h-8 rounded-full bg-emerald-100 text-emerald-600 font-bold text-sm">3</span>
              <div>
                <h3 className="font-semibold text-slate-800">Compare Documents</h3>
                <p className="text-xs text-slate-500">Click compare to find differences</p>
              </div>
            </div>
            
            <div className="pt-2 space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Compare Language</label>
                <div className="relative">
                  <GlobeAltIcon className="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-slate-400" />
                  <select className="w-full border border-slate-300 rounded-lg pl-10 pr-4 py-2.5 text-sm text-slate-700 appearance-none bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none">
                    <option>বাংলা (Bangla)</option>
                    <option>English</option>
                  </select>
                </div>
              </div>

              <button className="w-full bg-gradient-to-r from-indigo-500 to-blue-600 hover:from-indigo-600 hover:to-blue-700 text-white font-medium py-3 rounded-lg flex items-center justify-center space-x-2 shadow-md transition">
                <ArrowPathRoundedSquareIcon className="h-5 w-5" />
                <span>Compare Now</span>
              </button>
              
              <div className="flex items-start space-x-2 text-xs text-slate-500 mt-4">
                <CheckCircleIcon className="h-4 w-4 text-emerald-500 flex-shrink-0" />
                <p>Your files are secure and never stored permanently.</p>
              </div>
            </div>
          </div>
        </section>

        {/* Statistics Bar */}
        <section className="bg-white rounded-xl shadow-sm border border-slate-200 p-6 flex flex-wrap lg:flex-nowrap items-center justify-between gap-6">
          <div className="flex items-center space-x-6 pr-6 lg:border-r border-slate-200">
            <div className="relative h-24 w-24 flex items-center justify-center">
              <svg className="w-full h-full transform -rotate-90" viewBox="0 0 36 36">
                <path d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="#E2E8F0" strokeWidth="3" />
                <path d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="#10B981" strokeWidth="3" strokeDasharray="92, 100" />
              </svg>
              <div className="absolute text-2xl font-bold text-slate-800">92%</div>
            </div>
            <div>
              <h3 className="font-bold text-slate-800 mb-1">Overall Similarity</h3>
              <span className="inline-block bg-emerald-100 text-emerald-700 text-xs font-semibold px-2 py-1 rounded">Excellent Match</span>
              <p className="text-xs text-slate-500 mt-2">Documents are highly similar</p>
            </div>
          </div>

          <div className="flex-1 grid grid-cols-2 lg:grid-cols-4 gap-4">
            <div className="bg-red-50/50 p-4 rounded-xl border border-red-100 flex items-center space-x-4">
              <div className="bg-red-100 p-2 rounded-full text-red-500">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path></svg>
              </div>
              <div>
                <p className="text-xl font-bold text-slate-800">23</p>
                <p className="text-xs font-medium text-slate-600">Mismatched Words</p>
                <p className="text-[10px] text-slate-400">Found differences</p>
              </div>
            </div>
            
            <div className="bg-amber-50/50 p-4 rounded-xl border border-amber-100 flex items-center space-x-4">
              <div className="bg-amber-100 p-2 rounded-full text-amber-500">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M20 12H4"></path></svg>
              </div>
              <div>
                <p className="text-xl font-bold text-slate-800">7</p>
                <p className="text-xs font-medium text-slate-600">Missing Words</p>
                <p className="text-[10px] text-slate-400">Not in first document</p>
              </div>
            </div>

            <div className="bg-blue-50/50 p-4 rounded-xl border border-blue-100 flex items-center space-x-4">
              <div className="bg-blue-100 p-2 rounded-full text-blue-500">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 4v16m8-8H4"></path></svg>
              </div>
              <div>
                <p className="text-xl font-bold text-slate-800">5</p>
                <p className="text-xs font-medium text-slate-600">Extra Words</p>
                <p className="text-[10px] text-slate-400">Not in second document</p>
              </div>
            </div>

            <div className="bg-purple-50/50 p-4 rounded-xl border border-purple-100 flex items-center space-x-4">
              <div className="bg-purple-100 p-2 rounded-full text-purple-500">
                <DocumentTextIcon className="w-5 h-5" />
              </div>
              <div>
                <p className="text-xl font-bold text-slate-800">1,248</p>
                <p className="text-xs font-medium text-slate-600">Total Words Compared</p>
              </div>
            </div>
          </div>
        </section>

        {/* Side-by-side View */}
        <section className="grid grid-cols-1 md:grid-cols-2 gap-6 relative">
          {/* Middle sync icon */}
          <div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 bg-white border border-slate-200 rounded-full p-2 shadow-sm z-10 hidden md:block">
            <ArrowPathRoundedSquareIcon className="h-5 w-5 text-indigo-500" />
          </div>

          {/* Doc 1 Viewer */}
          <div className="bg-white rounded-xl shadow-sm border border-slate-200 flex flex-col h-[400px]">
            <div className="p-4 border-b border-slate-100 flex justify-between items-center bg-slate-50 rounded-t-xl">
              <div className="flex items-center space-x-2">
                <DocumentIcon className="h-5 w-5 text-red-500" />
                <h3 className="font-semibold text-slate-700 text-sm">First Document (Image/PDF)</h3>
              </div>
              <span className="text-xs text-slate-400">1,256 words</span>
            </div>
            <div className="p-6 overflow-y-auto flex-1 text-slate-700 leading-loose text-[15px] font-sans">
              বাংলাদেশ একটি <span className="bg-red-100 text-red-700 px-1 rounded border border-red-200">সুন্দর</span> দেশ। এই দেশের মানুষেরা খুব পরিশ্রমী ও <span className="bg-red-100 text-red-700 px-1 rounded border border-red-200">মেহনতি</span>। বাংলাদেশে অনেক প্রাকৃতিক <span className="bg-red-100 text-red-700 px-1 rounded border border-red-200">সৌন্দর্য</span> রয়েছে। পদ্মা, মেঘনা, যমুনা এই দেশের প্রধান নদী। আমাদের দেশকে আমরা ভালোবাসি এবং এর <span className="bg-red-100 text-red-700 px-1 rounded border border-red-200">উন্নতির</span> জন্য কাজ করি।
            </div>
          </div>

          {/* Doc 2 Viewer */}
          <div className="bg-white rounded-xl shadow-sm border border-slate-200 flex flex-col h-[400px]">
            <div className="p-4 border-b border-slate-100 flex justify-between items-center bg-slate-50 rounded-t-xl">
              <div className="flex items-center space-x-2">
                <DocumentTextIcon className="h-5 w-5 text-blue-500" />
                <h3 className="font-semibold text-slate-700 text-sm">Second Document (DOCX/PDF)</h3>
              </div>
              <span className="text-xs text-slate-400">1,261 words</span>
            </div>
            <div className="p-6 overflow-y-auto flex-1 text-slate-700 leading-loose text-[15px] font-sans">
              বাংলাদেশ একটি <span className="bg-emerald-100 text-emerald-800 px-1 rounded border border-emerald-200">সুন্দর</span> দেশ। এই দেশের মানুষেরা খুব পরিশ্রমী ও <span className="bg-emerald-100 text-emerald-800 px-1 rounded border border-emerald-200">মেহনতি</span>। বাংলাদেশে অনেক প্রাকৃতিক <span className="bg-emerald-100 text-emerald-800 px-1 rounded border border-emerald-200">সৌন্দর্য</span> রয়েছে। পদ্মা, মেঘনা, যমুনা এই দেশের প্রধান নদী। আমাদের দেশকে আমরা ভালোবাসি এবং এর <span className="bg-emerald-100 text-emerald-800 px-1 rounded border border-emerald-200">উন্নতির</span> জন্য কাজ করি।
            </div>
          </div>
        </section>

        {/* Differences Table and Export */}
        <section className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2 bg-white rounded-xl shadow-sm border border-slate-200 p-6">
            <h3 className="font-semibold text-slate-800 flex items-center space-x-2 mb-4">
              <span className="text-amber-500">⚠️</span>
              <span>Differences Found (23)</span>
            </h3>
            
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              {/* Diff Card 1 */}
              <div className="border border-slate-200 rounded-lg p-3 flex flex-col items-center justify-center space-y-2 relative">
                <div className="flex items-center space-x-3 w-full justify-center">
                  <span className="text-red-600 font-medium line-through">সুন্দর</span>
                  <span className="text-slate-400">→</span>
                  <span className="text-emerald-600 font-medium">সুন্দর</span>
                </div>
                <span className="text-[10px] text-slate-400 uppercase tracking-wider">Spelling</span>
              </div>
              
              {/* Diff Card 2 */}
              <div className="border border-slate-200 rounded-lg p-3 flex flex-col items-center justify-center space-y-2 relative">
                <div className="flex items-center space-x-3 w-full justify-center">
                  <span className="text-red-600 font-medium line-through">মেহনতি</span>
                  <span className="text-slate-400">→</span>
                  <span className="text-emerald-600 font-medium">মেহনতি</span>
                </div>
                <span className="text-[10px] text-slate-400 uppercase tracking-wider">Spelling</span>
              </div>

              {/* Diff Card 3 */}
              <div className="border border-slate-200 rounded-lg p-3 flex flex-col items-center justify-center space-y-2 relative">
                <div className="flex items-center space-x-3 w-full justify-center">
                  <span className="text-red-600 font-medium line-through">সৌন্দর্য</span>
                  <span className="text-slate-400">→</span>
                  <span className="text-emerald-600 font-medium">সৌন্দর্য</span>
                </div>
                <span className="text-[10px] text-slate-400 uppercase tracking-wider">Spelling</span>
              </div>

              {/* Diff Card 4 */}
              <div className="border border-slate-200 rounded-lg p-3 flex flex-col items-center justify-center space-y-2 relative">
                <div className="flex items-center space-x-3 w-full justify-center">
                  <span className="text-red-600 font-medium line-through">উন্নতির</span>
                  <span className="text-slate-400">→</span>
                  <span className="text-emerald-600 font-medium">উন্নতির</span>
                </div>
                <span className="text-[10px] text-slate-400 uppercase tracking-wider">Spelling</span>
              </div>
            </div>
          </div>

          <div className="bg-indigo-50/50 rounded-xl border border-indigo-100 p-6 flex flex-col justify-center">
            <div className="flex items-center space-x-3 mb-2">
              <ArrowDownTrayIcon className="h-5 w-5 text-indigo-600" />
              <h3 className="font-semibold text-slate-800">Export Report</h3>
            </div>
            <p className="text-xs text-slate-500 mb-6">Download comparison report</p>
            
            <div className="grid grid-cols-2 gap-3">
              <button className="flex items-center justify-center space-x-2 border border-red-200 hover:border-red-300 hover:bg-red-50 text-red-600 py-2.5 rounded-lg text-sm font-medium transition bg-white">
                <DocumentIcon className="h-4 w-4" />
                <span>PDF Report</span>
              </button>
              <button className="flex items-center justify-center space-x-2 border border-blue-200 hover:border-blue-300 hover:bg-blue-50 text-blue-600 py-2.5 rounded-lg text-sm font-medium transition bg-white">
                <DocumentTextIcon className="h-4 w-4" />
                <span>Word Report</span>
              </button>
            </div>
          </div>
        </section>

      </main>
    </div>
  );
}
