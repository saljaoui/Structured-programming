import { places } from "./where-do-we-go.data.js";
const sections = [];
let nameElement = null;
export const explore = () => {

  const sortedPlaces = sort()

  let namediv = document.createElement('div')
  namediv.className = "direction"
  document.body.append(namediv)


  console.log(sections);
  nameElement = document.createElement('a')
  nameElement.className = "location"
  nameElement.target = '_blank'
  nameElement.href = `https://www.google.com/maps/place/${places[0].coordinates}`
  nameElement.textContent = places[0].name + '\n' + places[0].coordinates
  nameElement.style.color = places[0].color
  document.body.append(nameElement)
let previousScrollY = null
  addEventListener("scroll", (event) => {
    const scrollY = window.scrollY;
    
    if (scrollY < (previousScrollY || 0)) {
      document.querySelector(".direction").innerHTML = "N";
    } else {
      document.querySelector(".direction").innerHTML = "S";
    }
 
    previousScrollY = scrollY;
    sections.forEach((section, index) => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.offsetHeight;
        if (scrollY >= sectionTop + (sectionHeight/2) - window.innerHeight) {
          if (nameElement) {
            nameElement.remove();
          }
               nameElement = document.createElement('a')
  nameElement.className = "location"
  nameElement.href = `https://www.google.com/maps/place/${places[index].coordinates}`
  nameElement.textContent = places[index].name + '\n' + places[index].coordinates
  nameElement.style.color = places[index].color
   nameElement.target = '_blank'
  
  document.body.append(nameElement)
        }
      });
  })

  sortedPlaces.forEach(place => {
      const section = document.createElement('section');
      const firstPart = place.name.split(",")[0].trim();
      section.style.backgroundImage = `url('${firstPart.toLowerCase().replace(/\s+/g, '-')}.jpg')`

      section.style.backgroundSize = 'cover';
      section.style.backgroundPosition = 'center';
      section.className = "section"
      section.style.display = 'flex';
      section.style.justifyContent = 'center';
      section.style.alignItems = 'center';
      sections.push(section);
      document.body.appendChild(section);
    });
  };

  const sort = () => {
    const northPlaces = places.filter(place => place.coordinates.includes("N"))
    const southPlaces = places.filter(place => place.coordinates.includes("S"))
    northPlaces.sort((a, b) => {
        if (a.coordinates > b.coordinates) return -1
        if (b.coordinates > a.coordinates) return 1   
    })
    southPlaces.sort((b, a) => {
        if (a.coordinates > b.coordinates) return -1
        if (b.coordinates > a.coordinates) return 1   
    })
    return northPlaces.concat(southPlaces)
}
