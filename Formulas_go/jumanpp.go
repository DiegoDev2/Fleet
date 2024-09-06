package main

// Jumanpp - Japanese Morphological Analyzer based on RNNLM
// Homepage: https://nlp.ist.i.kyoto-u.ac.jp/EN/index.php?JUMAN%2B%2B

import (
	"fmt"
	
	"os/exec"
)

func installJumanpp() {
	// Método 1: Descargar y extraer .tar.gz
	jumanpp_tar_url := "https://lotus.kuee.kyoto-u.ac.jp/nl-resource/jumanpp/jumanpp-1.02.tar.xz"
	jumanpp_cmd_tar := exec.Command("curl", "-L", jumanpp_tar_url, "-o", "package.tar.gz")
	err := jumanpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jumanpp_zip_url := "https://lotus.kuee.kyoto-u.ac.jp/nl-resource/jumanpp/jumanpp-1.02.tar.xz"
	jumanpp_cmd_zip := exec.Command("curl", "-L", jumanpp_zip_url, "-o", "package.zip")
	err = jumanpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jumanpp_bin_url := "https://lotus.kuee.kyoto-u.ac.jp/nl-resource/jumanpp/jumanpp-1.02.tar.xz"
	jumanpp_cmd_bin := exec.Command("curl", "-L", jumanpp_bin_url, "-o", "binary.bin")
	err = jumanpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jumanpp_src_url := "https://lotus.kuee.kyoto-u.ac.jp/nl-resource/jumanpp/jumanpp-1.02.tar.xz"
	jumanpp_cmd_src := exec.Command("curl", "-L", jumanpp_src_url, "-o", "source.tar.gz")
	err = jumanpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jumanpp_cmd_direct := exec.Command("./binary")
	err = jumanpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost-build")
exec.Command("latte", "install", "boost-build")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gperftools")
exec.Command("latte", "install", "gperftools")
}
