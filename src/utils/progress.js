export default class Progress {
    constructor () {
        this.count = 0;
        this.intervalCounterId = 0;
        this.progressbarEl = document.querySelector('#ngProgress');
    }

    start () {
        this.show();
        const self = this;
        clearInterval(this.intervalCounterId);
        this.intervalCounterId = setInterval(() => {
            if (isNaN(self.count)) {
                clearInterval(self.intervalCounterId);
                self.count = 0;
                self.hide();
            } else {
                self.remaining = 100 - self.count;
                self.count = self.count + (0.15 * Math.pow(1 - Math.sqrt(self.remaining), 2));
                self.updateCount(self.count);
            }
        }, 200);
    }

    /**
     * @param {int} new_count 
     */
    updateCount (new_count) {
        this.count = new_count;
        this.progressbarEl.style.width = `${this.count}%`;
    }

    hide () {
        this.progressbarEl.style.opacity = '0';
        const self = this;
        self.animate(() => {
            self.progressbarEl.style.width = '0%';
            self.animate(() => {
                self.show();
            }, 500);
        }, 500);
    }

    show () {
        const self = this;
        self.animate(() => {
            self.progressbarEl.style.opacity = '1';
        }, 100);
    }

    /**
     * Delay a given function by time
     * 
     * @param {Closure} fn 
     * @param {int} time 
     */
    animate (fn, time) {
        if (this.animation !== undefined) clearTimeout(this.animation);
        this.animation = setTimeout(fn, time);
    }

    stop () {
        clearInterval(this.intervalCounterId);
    }

    /**
     * @param {int} new_count 
     * @returns {int}
     */
    set (new_count) {
        this.show();
        this.updateCount(new_count);
        this.count = new_count;
        clearInterval(this.intervalCounterId);
        return this.count;
    }

    /**
     * @returns {int} 0
     */
    reset () {
        clearInterval(this.intervalCounterId);
        this.count = 0;
        this.updateCount(this.count);
        return 0;
    }

    /**
     * @returns {int}
     */
    complete () {
        this.count = 100;
        this.updateCount(this.count);
        const self = this;
        clearInterval(this.intervalCounterId);
        setTimeout(() => {
            self.hide();
            setTimeout(() => {
                self.count = 0;
                self.updateCount(self.count);
            }, 500);
        }, 1000);
        return this.count;
    }
}
