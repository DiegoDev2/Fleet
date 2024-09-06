package main

// Armadillo - C++ linear algebra library
// Homepage: https://arma.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installArmadillo() {
	// Método 1: Descargar y extraer .tar.gz
	armadillo_tar_url := "https://downloads.sourceforge.net/project/arma/armadillo-14.0.2.tar.xz"
	armadillo_cmd_tar := exec.Command("curl", "-L", armadillo_tar_url, "-o", "package.tar.gz")
	err := armadillo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	armadillo_zip_url := "https://downloads.sourceforge.net/project/arma/armadillo-14.0.2.tar.xz"
	armadillo_cmd_zip := exec.Command("curl", "-L", armadillo_zip_url, "-o", "package.zip")
	err = armadillo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	armadillo_bin_url := "https://downloads.sourceforge.net/project/arma/armadillo-14.0.2.tar.xz"
	armadillo_cmd_bin := exec.Command("curl", "-L", armadillo_bin_url, "-o", "binary.bin")
	err = armadillo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	armadillo_src_url := "https://downloads.sourceforge.net/project/arma/armadillo-14.0.2.tar.xz"
	armadillo_cmd_src := exec.Command("curl", "-L", armadillo_src_url, "-o", "source.tar.gz")
	err = armadillo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	armadillo_cmd_direct := exec.Command("./binary")
	err = armadillo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: arpack")
exec.Command("latte", "install", "arpack")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: superlu")
exec.Command("latte", "install", "superlu")
}
