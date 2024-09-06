package main

// R - Software environment for statistical computing
// Homepage: https://www.r-project.org/

import (
	"fmt"
	
	"os/exec"
)

func installR() {
	// Método 1: Descargar y extraer .tar.gz
	r_tar_url := "https://cran.r-project.org/src/base/R-4/R-4.4.1.tar.gz"
	r_cmd_tar := exec.Command("curl", "-L", r_tar_url, "-o", "package.tar.gz")
	err := r_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	r_zip_url := "https://cran.r-project.org/src/base/R-4/R-4.4.1.zip"
	r_cmd_zip := exec.Command("curl", "-L", r_zip_url, "-o", "package.zip")
	err = r_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	r_bin_url := "https://cran.r-project.org/src/base/R-4/R-4.4.1.bin"
	r_cmd_bin := exec.Command("curl", "-L", r_bin_url, "-o", "binary.bin")
	err = r_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	r_src_url := "https://cran.r-project.org/src/base/R-4/R-4.4.1.src.tar.gz"
	r_cmd_src := exec.Command("curl", "-L", r_src_url, "-o", "source.tar.gz")
	err = r_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	r_cmd_direct := exec.Command("./binary")
	err = r_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxau")
	exec.Command("latte", "install", "libxau").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxdmcp")
	exec.Command("latte", "install", "libxdmcp").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
