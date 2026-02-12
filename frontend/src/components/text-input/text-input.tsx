import React from "react";
import { Field } from "formik";

type TextInputProps = {
    name: string;
    label?: string;
    placeholder?: string;
    disabled?: boolean;
};

export const TextInput: React.FC<TextInputProps> = ({ name, label, placeholder, disabled }) => {
    return (
        <div className="text-input">
            {label && <label>{label}</label>}
            <Field name={name} placeholder={placeholder} disabled={disabled} />
        </div>
    );
};