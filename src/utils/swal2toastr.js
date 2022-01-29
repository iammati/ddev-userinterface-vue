import Sweetalert2 from "sweetalert2";

export default class Swal2Toastr {
    /**
     * Create a toast with sweetalert's mixins
     */
    createToastWrapper() {
        return Sweetalert2.mixin({
            toast: true,
            position: "bottom-end",
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            onOpen: (toast) => {
                toast.addEventListener("mouseenter", Sweetalert2.stopTimer);
                toast.addEventListener("mouseleave", Sweetalert2.resumeTimer);
            },
        });
    }
}
