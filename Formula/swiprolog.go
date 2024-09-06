package main

// SwiProlog - ISO/Edinburgh-style Prolog interpreter
// Homepage: https://www.swi-prolog.org/

import (
	"fmt"
	
	"os/exec"
)

func installSwiProlog() {
	// Método 1: Descargar y extraer .tar.gz
	swiprolog_tar_url := "https://www.swi-prolog.org/download/stable/src/swipl-9.2.7.tar.gz"
	swiprolog_cmd_tar := exec.Command("curl", "-L", swiprolog_tar_url, "-o", "package.tar.gz")
	err := swiprolog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiprolog_zip_url := "https://www.swi-prolog.org/download/stable/src/swipl-9.2.7.zip"
	swiprolog_cmd_zip := exec.Command("curl", "-L", swiprolog_zip_url, "-o", "package.zip")
	err = swiprolog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiprolog_bin_url := "https://www.swi-prolog.org/download/stable/src/swipl-9.2.7.bin"
	swiprolog_cmd_bin := exec.Command("curl", "-L", swiprolog_bin_url, "-o", "binary.bin")
	err = swiprolog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiprolog_src_url := "https://www.swi-prolog.org/download/stable/src/swipl-9.2.7.src.tar.gz"
	swiprolog_cmd_src := exec.Command("curl", "-L", swiprolog_src_url, "-o", "source.tar.gz")
	err = swiprolog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiprolog_cmd_direct := exec.Command("./binary")
	err = swiprolog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}
