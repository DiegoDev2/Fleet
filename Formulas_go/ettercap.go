package main

// Ettercap - Multipurpose sniffer/interceptor/logger for switched LAN
// Homepage: https://ettercap.github.io/ettercap/

import (
	"fmt"
	
	"os/exec"
)

func installEttercap() {
	// Método 1: Descargar y extraer .tar.gz
	ettercap_tar_url := "https://github.com/Ettercap/ettercap/archive/refs/tags/v0.8.3.1.tar.gz"
	ettercap_cmd_tar := exec.Command("curl", "-L", ettercap_tar_url, "-o", "package.tar.gz")
	err := ettercap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ettercap_zip_url := "https://github.com/Ettercap/ettercap/archive/refs/tags/v0.8.3.1.zip"
	ettercap_cmd_zip := exec.Command("curl", "-L", ettercap_zip_url, "-o", "package.zip")
	err = ettercap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ettercap_bin_url := "https://github.com/Ettercap/ettercap/archive/refs/tags/v0.8.3.1.bin"
	ettercap_cmd_bin := exec.Command("curl", "-L", ettercap_bin_url, "-o", "binary.bin")
	err = ettercap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ettercap_src_url := "https://github.com/Ettercap/ettercap/archive/refs/tags/v0.8.3.1.src.tar.gz"
	ettercap_cmd_src := exec.Command("curl", "-L", ettercap_src_url, "-o", "source.tar.gz")
	err = ettercap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ettercap_cmd_direct := exec.Command("./binary")
	err = ettercap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libmaxminddb")
exec.Command("latte", "install", "libmaxminddb")
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
