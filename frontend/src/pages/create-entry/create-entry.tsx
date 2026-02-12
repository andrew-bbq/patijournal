import React from "react";
import { Form, Formik } from "formik";
import { TextInput } from "../../components/text-input/text-input";

export const CreateEntry: React.FC = () => {
    return (
        <div className="entry">
            <Formik
                initialValues={{
                    title: '',
                    body: '',
                }}
                onSubmit={(values) => {console.log(values)}}
            >
                <Form>
                    <TextInput name='title' label='Title' placeholder='Entry Title' />
                    <TextInput name='body' label='Body' placeholder='Entry Body' />
                    <button type="submit">Create</button>
                </Form>
            </Formik>
        </div>
    );
};
