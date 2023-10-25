import React from "react";
import "./Infocard.css";

interface InfoCardProps {
  bgColor: string;
  title: string;
  count: number;
  icon: string;
}

const InfoCard = ({ bgColor, title, count, icon }: InfoCardProps) => {
  return (
    <div className={`info-box bg-${bgColor}-100  rounded-lg`}>
      <span className="info-icon --color-black">{icon}</span>
      <span className="info-text">
        <p>{title}</p>
        <h4>{count}</h4>
      </span>
    </div>
  );
};

export default InfoCard;