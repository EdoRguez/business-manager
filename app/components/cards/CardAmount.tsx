import { AmountType } from "@/app/common/enums";
import { IconType } from "react-icons";

interface cardAmountInterface {
  icon: IconType;
  iconColor: string;
  iconBackgroundColor: string;
  title: string;
  amount: number;
  amountType: AmountType;
}

const CardAmount: React.FC<cardAmountInterface> = ({
  icon: Icon,
  iconColor,
  iconBackgroundColor,
  title,
  amount,
  amountType
}) => {

  return (
    <div
      className="
        shadow-lg
        rounded-md 
        bg-white
        px-3 
        py-4"
    >
      <div
        className={`
          inline-flex
          p-1
          rounded
          ${iconBackgroundColor}
        `}
      >
        <Icon
          className={`
            inline-table
            ${iconColor}
          `}
          size={24}
        />
      </div>
      <div
        className="
          text-xs
          text-stone-700
          mt-3
        "
      >
        {title}
      </div>
      <div
        className="
          text-xl
          font-semibold
        "
      >
        { amountType === AmountType.Money ? "$" : "%" }
        { amount }
      </div>
    </div>
  );
};

export default CardAmount;
