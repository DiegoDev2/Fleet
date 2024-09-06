package main

// Calceph - C library to access the binary planetary ephemeris files
// Homepage: https://www.imcce.fr/inpop/calceph

import (
	"fmt"
	
	"os/exec"
)

func installCalceph() {
	// Método 1: Descargar y extraer .tar.gz
	calceph_tar_url := "https://www.imcce.fr/content/medias/recherche/equipes/asd/calceph/calceph-4.0.0.tar.gz"
	calceph_cmd_tar := exec.Command("curl", "-L", calceph_tar_url, "-o", "package.tar.gz")
	err := calceph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	calceph_zip_url := "https://www.imcce.fr/content/medias/recherche/equipes/asd/calceph/calceph-4.0.0.zip"
	calceph_cmd_zip := exec.Command("curl", "-L", calceph_zip_url, "-o", "package.zip")
	err = calceph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	calceph_bin_url := "https://www.imcce.fr/content/medias/recherche/equipes/asd/calceph/calceph-4.0.0.bin"
	calceph_cmd_bin := exec.Command("curl", "-L", calceph_bin_url, "-o", "binary.bin")
	err = calceph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	calceph_src_url := "https://www.imcce.fr/content/medias/recherche/equipes/asd/calceph/calceph-4.0.0.src.tar.gz"
	calceph_cmd_src := exec.Command("curl", "-L", calceph_src_url, "-o", "source.tar.gz")
	err = calceph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	calceph_cmd_direct := exec.Command("./binary")
	err = calceph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
