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
        px-4 
        py-5"
    >
      <div
        className={iconBackgroundColor}
      >
        <Icon
          className={iconColor}
          size={24}
        />
      </div>
      <div
        className="
          text-sm
        "
      >
        {title}
      </div>
      <div
        className="
          text-lg
        "
      >
        { amountType === AmountType.Money ? "$" : "%" }
        { amount }
      </div>
    </div>
  );
};

export default CardAmount;
