package main

// Singular - Computer algebra system for polynomial computations
// Homepage: https://www.singular.uni-kl.de/

import (
	"fmt"
	
	"os/exec"
)

func installSingular() {
	// Método 1: Descargar y extraer .tar.gz
	singular_tar_url := "https://www.singular.uni-kl.de/ftp/pub/Math/Singular/SOURCES/4-4-0/singular-4.4.0p5.tar.gz"
	singular_cmd_tar := exec.Command("curl", "-L", singular_tar_url, "-o", "package.tar.gz")
	err := singular_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	singular_zip_url := "https://www.singular.uni-kl.de/ftp/pub/Math/Singular/SOURCES/4-4-0/singular-4.4.0p5.zip"
	singular_cmd_zip := exec.Command("curl", "-L", singular_zip_url, "-o", "package.zip")
	err = singular_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	singular_bin_url := "https://www.singular.uni-kl.de/ftp/pub/Math/Singular/SOURCES/4-4-0/singular-4.4.0p5.bin"
	singular_cmd_bin := exec.Command("curl", "-L", singular_bin_url, "-o", "binary.bin")
	err = singular_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	singular_src_url := "https://www.singular.uni-kl.de/ftp/pub/Math/Singular/SOURCES/4-4-0/singular-4.4.0p5.src.tar.gz"
	singular_cmd_src := exec.Command("curl", "-L", singular_src_url, "-o", "source.tar.gz")
	err = singular_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	singular_cmd_direct := exec.Command("./binary")
	err = singular_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: ntl")
	exec.Command("latte", "install", "ntl").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
