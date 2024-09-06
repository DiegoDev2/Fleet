package main

// Libetonyek - Interpret and import Apple Keynote presentations
// Homepage: https://wiki.documentfoundation.org/DLP/Libraries/libetonyek

import (
	"fmt"
	
	"os/exec"
)

func installLibetonyek() {
	// Método 1: Descargar y extraer .tar.gz
	libetonyek_tar_url := "https://dev-www.libreoffice.org/src/libetonyek/libetonyek-0.1.10.tar.xz"
	libetonyek_cmd_tar := exec.Command("curl", "-L", libetonyek_tar_url, "-o", "package.tar.gz")
	err := libetonyek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libetonyek_zip_url := "https://dev-www.libreoffice.org/src/libetonyek/libetonyek-0.1.10.tar.xz"
	libetonyek_cmd_zip := exec.Command("curl", "-L", libetonyek_zip_url, "-o", "package.zip")
	err = libetonyek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libetonyek_bin_url := "https://dev-www.libreoffice.org/src/libetonyek/libetonyek-0.1.10.tar.xz"
	libetonyek_cmd_bin := exec.Command("curl", "-L", libetonyek_bin_url, "-o", "binary.bin")
	err = libetonyek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libetonyek_src_url := "https://dev-www.libreoffice.org/src/libetonyek/libetonyek-0.1.10.tar.xz"
	libetonyek_cmd_src := exec.Command("curl", "-L", libetonyek_src_url, "-o", "source.tar.gz")
	err = libetonyek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libetonyek_cmd_direct := exec.Command("./binary")
	err = libetonyek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: glm")
	exec.Command("latte", "install", "glm").Run()
	fmt.Println("Instalando dependencia: mdds")
	exec.Command("latte", "install", "mdds").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
}
