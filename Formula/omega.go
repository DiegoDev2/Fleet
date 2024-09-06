package main

// Omega - Packaged search engine for websites, built on top of Xapian
// Homepage: https://xapian.org/

import (
	"fmt"
	
	"os/exec"
)

func installOmega() {
	// Método 1: Descargar y extraer .tar.gz
	omega_tar_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-omega-1.4.26.tar.xz"
	omega_cmd_tar := exec.Command("curl", "-L", omega_tar_url, "-o", "package.tar.gz")
	err := omega_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	omega_zip_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-omega-1.4.26.tar.xz"
	omega_cmd_zip := exec.Command("curl", "-L", omega_zip_url, "-o", "package.zip")
	err = omega_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	omega_bin_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-omega-1.4.26.tar.xz"
	omega_cmd_bin := exec.Command("curl", "-L", omega_bin_url, "-o", "binary.bin")
	err = omega_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	omega_src_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-omega-1.4.26.tar.xz"
	omega_cmd_src := exec.Command("curl", "-L", omega_src_url, "-o", "source.tar.gz")
	err = omega_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	omega_cmd_direct := exec.Command("./binary")
	err = omega_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: xapian")
	exec.Command("latte", "install", "xapian").Run()
}
