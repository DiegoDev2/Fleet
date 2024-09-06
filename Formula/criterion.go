package main

// Criterion - Cross-platform C and C++ unit testing framework for the 21st century
// Homepage: https://github.com/Snaipe/Criterion

import (
	"fmt"
	
	"os/exec"
)

func installCriterion() {
	// Método 1: Descargar y extraer .tar.gz
	criterion_tar_url := "https://github.com/Snaipe/Criterion/releases/download/v2.4.2/criterion-2.4.2.tar.xz"
	criterion_cmd_tar := exec.Command("curl", "-L", criterion_tar_url, "-o", "package.tar.gz")
	err := criterion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	criterion_zip_url := "https://github.com/Snaipe/Criterion/releases/download/v2.4.2/criterion-2.4.2.tar.xz"
	criterion_cmd_zip := exec.Command("curl", "-L", criterion_zip_url, "-o", "package.zip")
	err = criterion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	criterion_bin_url := "https://github.com/Snaipe/Criterion/releases/download/v2.4.2/criterion-2.4.2.tar.xz"
	criterion_cmd_bin := exec.Command("curl", "-L", criterion_bin_url, "-o", "binary.bin")
	err = criterion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	criterion_src_url := "https://github.com/Snaipe/Criterion/releases/download/v2.4.2/criterion-2.4.2.tar.xz"
	criterion_cmd_src := exec.Command("curl", "-L", criterion_src_url, "-o", "source.tar.gz")
	err = criterion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	criterion_cmd_direct := exec.Command("./binary")
	err = criterion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: nanomsg")
	exec.Command("latte", "install", "nanomsg").Run()
	fmt.Println("Instalando dependencia: nanopb")
	exec.Command("latte", "install", "nanopb").Run()
}
