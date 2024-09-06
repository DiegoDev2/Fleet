package main

// Pdfpc - Presenter console with multi-monitor support for PDF files
// Homepage: https://pdfpc.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installPdfpc() {
	// Método 1: Descargar y extraer .tar.gz
	pdfpc_tar_url := "https://github.com/pdfpc/pdfpc/archive/refs/tags/v4.6.0.tar.gz"
	pdfpc_cmd_tar := exec.Command("curl", "-L", pdfpc_tar_url, "-o", "package.tar.gz")
	err := pdfpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfpc_zip_url := "https://github.com/pdfpc/pdfpc/archive/refs/tags/v4.6.0.zip"
	pdfpc_cmd_zip := exec.Command("curl", "-L", pdfpc_zip_url, "-o", "package.zip")
	err = pdfpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfpc_bin_url := "https://github.com/pdfpc/pdfpc/archive/refs/tags/v4.6.0.bin"
	pdfpc_cmd_bin := exec.Command("curl", "-L", pdfpc_bin_url, "-o", "binary.bin")
	err = pdfpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfpc_src_url := "https://github.com/pdfpc/pdfpc/archive/refs/tags/v4.6.0.src.tar.gz"
	pdfpc_cmd_src := exec.Command("curl", "-L", pdfpc_src_url, "-o", "source.tar.gz")
	err = pdfpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfpc_cmd_direct := exec.Command("./binary")
	err = pdfpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: discount")
	exec.Command("latte", "install", "discount").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gstreamer")
	exec.Command("latte", "install", "gstreamer").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libgee")
	exec.Command("latte", "install", "libgee").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: webkitgtk")
	exec.Command("latte", "install", "webkitgtk").Run()
}
