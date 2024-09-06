package main

// Abyss - Genome sequence assembler for short reads
// Homepage: https://www.bcgsc.ca/resources/software/abyss

import (
	"fmt"
	
	"os/exec"
)

func installAbyss() {
	// Método 1: Descargar y extraer .tar.gz
	abyss_tar_url := "https://github.com/bcgsc/abyss/releases/download/2.3.8/abyss-2.3.8.tar.gz"
	abyss_cmd_tar := exec.Command("curl", "-L", abyss_tar_url, "-o", "package.tar.gz")
	err := abyss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abyss_zip_url := "https://github.com/bcgsc/abyss/releases/download/2.3.8/abyss-2.3.8.zip"
	abyss_cmd_zip := exec.Command("curl", "-L", abyss_zip_url, "-o", "package.zip")
	err = abyss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abyss_bin_url := "https://github.com/bcgsc/abyss/releases/download/2.3.8/abyss-2.3.8.bin"
	abyss_cmd_bin := exec.Command("curl", "-L", abyss_bin_url, "-o", "binary.bin")
	err = abyss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abyss_src_url := "https://github.com/bcgsc/abyss/releases/download/2.3.8/abyss-2.3.8.src.tar.gz"
	abyss_cmd_src := exec.Command("curl", "-L", abyss_src_url, "-o", "source.tar.gz")
	err = abyss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abyss_cmd_direct := exec.Command("./binary")
	err = abyss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: multimarkdown")
exec.Command("latte", "install", "multimarkdown")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: google-sparsehash")
exec.Command("latte", "install", "google-sparsehash")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
