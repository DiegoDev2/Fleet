package main

// Wireshark - Network analyzer and capture tool - without graphical user interface
// Homepage: https://www.wireshark.org

import (
	"fmt"
	
	"os/exec"
)

func installWireshark() {
	// Método 1: Descargar y extraer .tar.gz
	wireshark_tar_url := "https://www.wireshark.org/download/src/all-versions/wireshark-4.4.0.tar.xz"
	wireshark_cmd_tar := exec.Command("curl", "-L", wireshark_tar_url, "-o", "package.tar.gz")
	err := wireshark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wireshark_zip_url := "https://www.wireshark.org/download/src/all-versions/wireshark-4.4.0.tar.xz"
	wireshark_cmd_zip := exec.Command("curl", "-L", wireshark_zip_url, "-o", "package.zip")
	err = wireshark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wireshark_bin_url := "https://www.wireshark.org/download/src/all-versions/wireshark-4.4.0.tar.xz"
	wireshark_cmd_bin := exec.Command("curl", "-L", wireshark_bin_url, "-o", "binary.bin")
	err = wireshark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wireshark_src_url := "https://www.wireshark.org/download/src/all-versions/wireshark-4.4.0.tar.xz"
	wireshark_cmd_src := exec.Command("curl", "-L", wireshark_src_url, "-o", "source.tar.gz")
	err = wireshark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wireshark_cmd_direct := exec.Command("./binary")
	err = wireshark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: c-ares")
	exec.Command("latte", "install", "c-ares").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libmaxminddb")
	exec.Command("latte", "install", "libmaxminddb").Run()
	fmt.Println("Instalando dependencia: libnghttp2")
	exec.Command("latte", "install", "libnghttp2").Run()
	fmt.Println("Instalando dependencia: libnghttp3")
	exec.Command("latte", "install", "libnghttp3").Run()
	fmt.Println("Instalando dependencia: libsmi")
	exec.Command("latte", "install", "libsmi").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: speexdsp")
	exec.Command("latte", "install", "speexdsp").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
