package main

// CargoAllFeatures - Cargo subcommands to build and test all feature flag combinations
// Homepage: https://github.com/frewsxcv/cargo-all-features

import (
	"fmt"
	
	"os/exec"
)

func installCargoAllFeatures() {
	// Método 1: Descargar y extraer .tar.gz
	cargoallfeatures_tar_url := "https://github.com/frewsxcv/cargo-all-features/archive/refs/tags/1.10.0.tar.gz"
	cargoallfeatures_cmd_tar := exec.Command("curl", "-L", cargoallfeatures_tar_url, "-o", "package.tar.gz")
	err := cargoallfeatures_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoallfeatures_zip_url := "https://github.com/frewsxcv/cargo-all-features/archive/refs/tags/1.10.0.zip"
	cargoallfeatures_cmd_zip := exec.Command("curl", "-L", cargoallfeatures_zip_url, "-o", "package.zip")
	err = cargoallfeatures_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoallfeatures_bin_url := "https://github.com/frewsxcv/cargo-all-features/archive/refs/tags/1.10.0.bin"
	cargoallfeatures_cmd_bin := exec.Command("curl", "-L", cargoallfeatures_bin_url, "-o", "binary.bin")
	err = cargoallfeatures_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoallfeatures_src_url := "https://github.com/frewsxcv/cargo-all-features/archive/refs/tags/1.10.0.src.tar.gz"
	cargoallfeatures_cmd_src := exec.Command("curl", "-L", cargoallfeatures_src_url, "-o", "source.tar.gz")
	err = cargoallfeatures_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoallfeatures_cmd_direct := exec.Command("./binary")
	err = cargoallfeatures_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
