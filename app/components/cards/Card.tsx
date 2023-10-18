import { AiOutlineInfoCircle } from "react-icons/ai";
import { HiArrowLongRight } from "react-icons/hi2";
import { Swiper, SwiperSlide } from "swiper/react";

import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";

import { Autoplay, Pagination, Navigation } from "swiper/modules";
import Button from "../Button";

const Card = () => {
  return (
    <div
      className="
        shadow-lg
        rounded-md 
        bg-white
        px-4 
        py-5
        transition 
        duration-200
        hover:scale-105"
    >
      <div
        className="
            rounded-md 
            bg-zinc-50"
      >
        <Swiper
          spaceBetween={30}
          centeredSlides={true}
          grabCursor={true}
          pagination={{
            enabled: true,
            clickable: true,
          }}
          modules={[Pagination]}
        >
          <SwiperSlide>
            <div className="bg-center bg-no-repeat bg-contain bg-[url('https://swiperjs.com/demos/images/nature-1.jpg')] h-60 w-full"></div>
          </SwiperSlide>
          <SwiperSlide>
            <div className="bg-center bg-no-repeat bg-contain bg-[url('https://swiperjs.com/demos/images/nature-2.jpg')] h-60 w-full"></div>
          </SwiperSlide>
          <SwiperSlide>
            <div className="bg-center bg-no-repeat bg-contain bg-[url('https://swiperjs.com/demos/images/nature-3.jpg')] h-60 w-full"></div>
          </SwiperSlide>
          <SwiperSlide>
            <div className="bg-center bg-no-repeat bg-contain bg-[url('https://swiperjs.com/demos/images/nature-4.jpg')] h-60 w-full"></div>
          </SwiperSlide>
        </Swiper>
      </div>
      <div
        className="
            flex 
            flex-col"
      >
        <span
          className="
                text-center 
                text-lg 
                font-bold 
                uppercase my-4"
        >
          Pedido 1
        </span>
        <table
          className="
                table-auto"
        >
          <tbody>
            <tr>
              <td>
                <small>Fecha:</small>
              </td>
              <td className="text-right">
                <small>16/12/2023</small>
              </td>
            </tr>
            <tr>
              <td>
                <small>Ganancias:</small>
              </td>
              <td className="text-right">
                <small>100.00 USD</small>
              </td>
            </tr>
          </tbody>
        </table>
        <div className="mt-8">
          <Button
            label={"Detalles"}
            icon={AiOutlineInfoCircle}
            onClick={() => {}}
          />
        </div>
      </div>
    </div>
  );
};

export default Card;
