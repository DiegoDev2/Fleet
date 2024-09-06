package main

// W3m - Pager/text based browser
// Homepage: https://w3m.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installW3m() {
	// Método 1: Descargar y extraer .tar.gz
	w3m_tar_url := "https://deb.debian.org/debian/pool/main/w/w3m/w3m_0.5.3.orig.tar.gz"
	w3m_cmd_tar := exec.Command("curl", "-L", w3m_tar_url, "-o", "package.tar.gz")
	err := w3m_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	w3m_zip_url := "https://deb.debian.org/debian/pool/main/w/w3m/w3m_0.5.3.orig.zip"
	w3m_cmd_zip := exec.Command("curl", "-L", w3m_zip_url, "-o", "package.zip")
	err = w3m_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	w3m_bin_url := "https://deb.debian.org/debian/pool/main/w/w3m/w3m_0.5.3.orig.bin"
	w3m_cmd_bin := exec.Command("curl", "-L", w3m_bin_url, "-o", "binary.bin")
	err = w3m_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	w3m_src_url := "https://deb.debian.org/debian/pool/main/w/w3m/w3m_0.5.3.orig.src.tar.gz"
	w3m_cmd_src := exec.Command("curl", "-L", w3m_src_url, "-o", "source.tar.gz")
	err = w3m_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	w3m_cmd_direct := exec.Command("./binary")
	err = w3m_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libbsd")
	exec.Command("latte", "install", "libbsd").Run()
}
