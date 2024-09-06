package main

// Cherrytree - Hierarchical note taking application featuring rich text and syntax highlighting
// Homepage: https://www.giuspen.com/cherrytree/

import (
	"fmt"
	
	"os/exec"
)

func installCherrytree() {
	// Método 1: Descargar y extraer .tar.gz
	cherrytree_tar_url := "https://www.giuspen.com/software/cherrytree_1.1.4.tar.xz"
	cherrytree_cmd_tar := exec.Command("curl", "-L", cherrytree_tar_url, "-o", "package.tar.gz")
	err := cherrytree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cherrytree_zip_url := "https://www.giuspen.com/software/cherrytree_1.1.4.tar.xz"
	cherrytree_cmd_zip := exec.Command("curl", "-L", cherrytree_zip_url, "-o", "package.zip")
	err = cherrytree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cherrytree_bin_url := "https://www.giuspen.com/software/cherrytree_1.1.4.tar.xz"
	cherrytree_cmd_bin := exec.Command("curl", "-L", cherrytree_bin_url, "-o", "binary.bin")
	err = cherrytree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cherrytree_src_url := "https://www.giuspen.com/software/cherrytree_1.1.4.tar.xz"
	cherrytree_cmd_src := exec.Command("curl", "-L", cherrytree_src_url, "-o", "source.tar.gz")
	err = cherrytree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cherrytree_cmd_direct := exec.Command("./binary")
	err = cherrytree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: atkmm@2.28")
exec.Command("latte", "install", "atkmm@2.28")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: cairomm@1.14")
exec.Command("latte", "install", "cairomm@1.14")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: fribidi")
exec.Command("latte", "install", "fribidi")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glibmm@2.66")
exec.Command("latte", "install", "glibmm@2.66")
	fmt.Println("Instalando dependencia: gspell")
exec.Command("latte", "install", "gspell")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: gtkmm3")
exec.Command("latte", "install", "gtkmm3")
	fmt.Println("Instalando dependencia: gtksourceview3")
exec.Command("latte", "install", "gtksourceview3")
	fmt.Println("Instalando dependencia: gtksourceviewmm3")
exec.Command("latte", "install", "gtksourceviewmm3")
	fmt.Println("Instalando dependencia: libsigc++@2")
exec.Command("latte", "install", "libsigc++@2")
	fmt.Println("Instalando dependencia: libxml++")
exec.Command("latte", "install", "libxml++")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: pangomm@2.46")
exec.Command("latte", "install", "pangomm@2.46")
	fmt.Println("Instalando dependencia: spdlog")
exec.Command("latte", "install", "spdlog")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: uchardet")
exec.Command("latte", "install", "uchardet")
	fmt.Println("Instalando dependencia: vte3")
exec.Command("latte", "install", "vte3")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: enchant")
exec.Command("latte", "install", "enchant")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
