export type Result = {
  status: 'LOADING'|'OK'|'ERROR',
  action: string
}

export const MAX_INT = 2147483647;

export const NUMERIC_REGEX = /[0-9]/;
export const LOWER_CASE_REGEX = /[a-zšđžčć]/;
export const UPPER_CASE_REGEX = /[A-ZŠĐŽČĆ]/;
export const ALPHA_REGEX = /^[a-zA-ZčČćĆđĐžŽšŠ ]+$/;
export const ALPHANUMERIC_REGEX = /^[a-zA-ZčČćĆđĐžŽšŠ0-9 ]+$/;
export const SPECIAL_CHARACTERS_REGEX = /[^a-zA-Z0-9]/;
export const EMAIL_REGEX = /^([\w-.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$/;
export const SCRIPT_REGEX = /^(?!.*(\<script\>|\<\/script\>|\&lt\;script\&gt\;|\&lt\;\/script\&gt\;|javascript\:)).*/;
export const PHONE_NUMBER_REGEX = /^(\+\d{1,3})?\-?\(?\d{2,3}\)?[-]?\d{3}[-]?\d{3,4}$/;