package main

// Julia - Fast, Dynamic Programming Language
// Homepage: https://julialang.org/

import (
	"fmt"
	
	"os/exec"
)

func installJulia() {
	// Método 1: Descargar y extraer .tar.gz
	julia_tar_url := "https://github.com/JuliaLang/julia/releases/download/v1.10.4/julia-1.10.4-full.tar.gz"
	julia_cmd_tar := exec.Command("curl", "-L", julia_tar_url, "-o", "package.tar.gz")
	err := julia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	julia_zip_url := "https://github.com/JuliaLang/julia/releases/download/v1.10.4/julia-1.10.4-full.zip"
	julia_cmd_zip := exec.Command("curl", "-L", julia_zip_url, "-o", "package.zip")
	err = julia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	julia_bin_url := "https://github.com/JuliaLang/julia/releases/download/v1.10.4/julia-1.10.4-full.bin"
	julia_cmd_bin := exec.Command("curl", "-L", julia_bin_url, "-o", "binary.bin")
	err = julia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	julia_src_url := "https://github.com/JuliaLang/julia/releases/download/v1.10.4/julia-1.10.4-full.src.tar.gz"
	julia_cmd_src := exec.Command("curl", "-L", julia_src_url, "-o", "source.tar.gz")
	err = julia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	julia_cmd_direct := exec.Command("./binary")
	err = julia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libblastrampoline")
	exec.Command("latte", "install", "libblastrampoline").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: libnghttp2")
	exec.Command("latte", "install", "libnghttp2").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: mbedtls@2")
	exec.Command("latte", "install", "mbedtls@2").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: openlibm")
	exec.Command("latte", "install", "openlibm").Run()
	fmt.Println("Instalando dependencia: p7zip")
	exec.Command("latte", "install", "p7zip").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: utf8proc")
	exec.Command("latte", "install", "utf8proc").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
