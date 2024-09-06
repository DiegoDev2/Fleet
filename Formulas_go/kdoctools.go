package main

// Kdoctools - Create documentation from DocBook
// Homepage: https://api.kde.org/frameworks/kdoctools/html/index.html

import (
	"fmt"
	
	"os/exec"
)

func installKdoctools() {
	// Método 1: Descargar y extraer .tar.gz
	kdoctools_tar_url := "https://download.kde.org/stable/frameworks/6.5/kdoctools-6.5.0.tar.xz"
	kdoctools_cmd_tar := exec.Command("curl", "-L", kdoctools_tar_url, "-o", "package.tar.gz")
	err := kdoctools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kdoctools_zip_url := "https://download.kde.org/stable/frameworks/6.5/kdoctools-6.5.0.tar.xz"
	kdoctools_cmd_zip := exec.Command("curl", "-L", kdoctools_zip_url, "-o", "package.zip")
	err = kdoctools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kdoctools_bin_url := "https://download.kde.org/stable/frameworks/6.5/kdoctools-6.5.0.tar.xz"
	kdoctools_cmd_bin := exec.Command("curl", "-L", kdoctools_bin_url, "-o", "binary.bin")
	err = kdoctools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kdoctools_src_url := "https://download.kde.org/stable/frameworks/6.5/kdoctools-6.5.0.tar.xz"
	kdoctools_cmd_src := exec.Command("curl", "-L", kdoctools_src_url, "-o", "source.tar.gz")
	err = kdoctools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kdoctools_cmd_direct := exec.Command("./binary")
	err = kdoctools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: extra-cmake-modules")
exec.Command("latte", "install", "extra-cmake-modules")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: ki18n")
exec.Command("latte", "install", "ki18n")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: karchive")
exec.Command("latte", "install", "karchive")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
